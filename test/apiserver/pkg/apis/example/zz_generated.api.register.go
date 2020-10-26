// Code generated by apiregister-gen. DO NOT EDIT.

package example

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-builder-alpha/pkg/builders"
)

var (
	ExampleContactStorage = builders.NewApiResource( // Resource status endpoint
		InternalContact,
		func() runtime.Object { return &Contact{} },     // Register versioned resource
		func() runtime.Object { return &ContactList{} }, // Register versioned resource list
		&ContactStrategy{builders.StorageStrategySingleton},
	)
	InternalContact = builders.NewInternalResource(
		"contacts",
		"Contact",
		func() runtime.Object { return &Contact{} },
		func() runtime.Object { return &ContactList{} },
	)
	InternalContactStatus = builders.NewInternalResourceStatus(
		"contacts",
		"ContactStatus",
		func() runtime.Object { return &Contact{} },
		func() runtime.Object { return &ContactList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("example.infoblox.com").WithKinds(
		InternalContact,
		InternalContactStatus,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
	}).AddToScheme
	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Contact struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ContactSpec
	Status ContactStatus
}

type ContactSpec struct {
}

type ContactStatus struct {
}

//
// Contact Functions and Structs
//
// +k8s:deepcopy-gen=false
type ContactStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ContactStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ContactList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Contact
}

func (Contact) NewStatus() interface{} {
	return ContactStatus{}
}

func (pc *Contact) GetStatus() interface{} {
	return pc.Status
}

func (pc *Contact) SetStatus(s interface{}) {
	pc.Status = s.(ContactStatus)
}

func (pc *Contact) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Contact) SetSpec(s interface{}) {
	pc.Spec = s.(ContactSpec)
}

func (pc *Contact) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Contact) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Contact) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Contact.
// +k8s:deepcopy-gen=false
type ContactRegistry interface {
	ListContacts(ctx context.Context, options *internalversion.ListOptions) (*ContactList, error)
	GetContact(ctx context.Context, id string, options *metav1.GetOptions) (*Contact, error)
	CreateContact(ctx context.Context, id *Contact) (*Contact, error)
	UpdateContact(ctx context.Context, id *Contact) (*Contact, error)
	DeleteContact(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewContactRegistry(sp builders.StandardStorageProvider) ContactRegistry {
	return &storageContact{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageContact struct {
	builders.StandardStorageProvider
}

func (s *storageContact) ListContacts(ctx context.Context, options *internalversion.ListOptions) (*ContactList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ContactList), err
}

func (s *storageContact) GetContact(ctx context.Context, id string, options *metav1.GetOptions) (*Contact, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Contact), nil
}

func (s *storageContact) CreateContact(ctx context.Context, object *Contact) (*Contact, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Contact), nil
}

func (s *storageContact) UpdateContact(ctx context.Context, object *Contact) (*Contact, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*Contact), nil
}

func (s *storageContact) DeleteContact(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil, &metav1.DeleteOptions{})
	return sync, err
}
