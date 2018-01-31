package controller

import (
	clientset "crds/pkg/client/clientset/versioned"
	crdinformer "crds/pkg/client/informers/externalversions"
	configinformer "crds/pkg/client/informers/externalversions/dev/v1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"

	"path/filepath"

	v1 "crds/pkg/apis/dev.kubernetes.pavanvelidandla.com/v1"
	git "gopkg.in/src-d/go-git.v4"
	k8v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	crdclientset clientset.Interface
	queue        workqueue.RateLimitingInterface
	informer     configinformer.ConfigFromGitInformer
	kubeclient   kubernetes.Interface
}

func NewController(crdclientset clientset.Interface, kubeclient kubernetes.Interface) *Controller {

	crdinfromersfactory := crdinformer.NewSharedInformerFactory(crdclientset, time.Second*30)
	configsfromgitinformer := crdinfromersfactory.Dev().V1().ConfigFromGits()
	configsqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	log.Println("Setting up event handlers")

	configsfromgitinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				configsqueue.Add(key)
			}
		},

		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*v1.ConfigFromGit)
			oldDepl := old.(*v1.ConfigFromGit)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Resources.
				// Two different versions of the same CRDS (configgit kind) will always have different RVs.
				return
			}

			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				configsqueue.Add(key)
			}
		},
	})

	controller := &Controller{
		crdclientset: crdclientset,
		informer:     configsfromgitinformer,
		queue:        configsqueue,
		kubeclient:   kubeclient,
	}

	return controller

}

func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	log.Println("Starting kubewatch controller")

	go c.informer.Informer().Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	log.Println("Kubewatch controller synced and ready")

	wait.Until(c.runWorker, time.Second, stopCh)
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func (c *Controller) HasSynced() bool {
	return c.informer.Informer().HasSynced()
}

// LastSyncResourceVersion is required for the cache.Controller interface.
func (c *Controller) LastSyncResourceVersion() string {
	return c.informer.Informer().LastSyncResourceVersion()
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < 5 {
		log.Println("Error processing %s (will retry): %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		log.Println("Error processing %s (giving up): %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}

	return true
}

func (c *Controller) processItem(key string) error {
	//c.logger.Infof("Processing change to Pod %s", key)

	Obj, exists, err := c.informer.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}

	if !exists {
		//fmt.Printf(obj)
		//c.eventHandler.ObjectDeleted(obj)
		//fmt.Print(obj)
		fmt.Println("Object terminated - ", key)
		return nil
	}
	//fmt.Print(obj)
	//fmt.Println("Created a new Pod ", key, Obj.(*api_v1.Pod).Name, " Container Name - ", Obj.(*api_v1.Pod).Spec.Containers[0].Name, " Image Name - ", Obj.(*api_v1.Pod).Spec.Containers[0].Image)
	fmt.Println("Created a new crd ", Obj.(*v1.ConfigFromGit).Name, " ", Obj.(*v1.ConfigFromGit).Spec.GitUrl)
	c.ProcessConfig(*Obj.(*v1.ConfigFromGit))
	return nil
}

func (c *Controller) ProcessConfig(Obj v1.ConfigFromGit) {

	log.Println("Processing request for ", Obj.Name, " ", Obj.Spec.GitUrl)

	GitUrl := Obj.Spec.GitUrl
	directory := "/tmp/gitworkspace/" + Obj.GetObjectMeta().GetResourceVersion() + "/*"
	log.Println("Directory Path - ", directory)
	os.RemoveAll(directory)
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               GitUrl,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		os.RemoveAll(directory)
		log.Println("Failed to checkout ", err)
		return
	}
	log.Println("Removing .git folder ")
	os.RemoveAll(directory + ".git")
	Data := make(map[string]string)

	_ = filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {

		b, err := ioutil.ReadFile(path)
		filecontent := string(b)
		Data[filepath.Base(path)] = filecontent
		return nil
	})

	cm := &k8v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      Obj.Name,
			Namespace: Obj.Namespace,
		},

		Data: Data,
	}
	result, err := c.kubeclient.CoreV1().ConfigMaps("default").Create(cm)
	if err != nil {
		log.Println("Error creating configmap ", err)
		os.RemoveAll(directory)
		return
	}

	log.Println("Created configmap from Git", result.GetObjectMeta().GetName())
	os.RemoveAll(directory)
}
