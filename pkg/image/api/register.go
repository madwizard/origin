package api

import (
	"k8s.io/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&Image{},
		&ImageList{},
		&ImageStream{},
		&ImageStreamList{},
		&ImageStreamMapping{},
		&ImageStreamTag{},
		&ImageStreamTagList{},
		&ImageStreamImage{},
		&DockerImage{},
	)
}

func (*Image) IsAnAPIObject()              {}
func (*ImageList) IsAnAPIObject()          {}
func (*DockerImage) IsAnAPIObject()        {}
func (*ImageStream) IsAnAPIObject()        {}
func (*ImageStreamList) IsAnAPIObject()    {}
func (*ImageStreamMapping) IsAnAPIObject() {}
func (*ImageStreamTag) IsAnAPIObject()     {}
func (*ImageStreamTagList) IsAnAPIObject() {}
func (*ImageStreamImage) IsAnAPIObject()   {}
