package controllers

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

type cacheGetter interface {
	getCache(name string) cache.Indexer
}

func generateUpdate(c cacheGetter, name string) func(event.UpdateEvent, workqueue.RateLimitingInterface) {
	indexer := c.getCache(name)
	return func(updateEvent event.UpdateEvent, limitingInterface workqueue.RateLimitingInterface) {
		//if updateEvent.ObjectOld != nil {
		//	if err := indexer.Delete(updateEvent.ObjectOld); err != nil {
		//		//todo
		//	}
		//} else {
		//}

		if updateEvent.ObjectNew != nil {
			if err := indexer.Update(updateEvent.ObjectNew); err != nil {
				//todo
			}
		} else {
		}
	}
}

func generateCreate(c cacheGetter, name string) func(event.CreateEvent, workqueue.RateLimitingInterface) {
	indexer := c.getCache(name)
	return func(createEvent event.CreateEvent, limitingInterface workqueue.RateLimitingInterface) {
		if createEvent.Object != nil {
			if err := indexer.Add(createEvent); err != nil {
				return
			}
		}
	}
}

func generateDelete(c cacheGetter, name string) func(event.DeleteEvent, workqueue.RateLimitingInterface) {
	indexer := c.getCache(name)
	return func(deleteEvent event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
		if deleteEvent.Object != nil {
			if err := indexer.Delete(deleteEvent.Object); err != nil {
				return
			}
		}
	}
}

func generalHandler(c cacheGetter, name string) *handler.Funcs {
	return &handler.Funcs{
		CreateFunc: generateCreate(c, name),
		UpdateFunc: generateUpdate(c, name),
		DeleteFunc: generateDelete(c, name),
	}
}
