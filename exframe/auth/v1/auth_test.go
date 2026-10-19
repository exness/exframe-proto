package authv1

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestExtension(t *testing.T) {
	t.Parallel()

	ext := E_Method.TypeDescriptor()
	if got := ext.Number(); got != 52008 {
		t.Fatalf("extension number = %d, want 52008", got)
	}
	if got := ext.FullName(); got != "exframe.auth.v1.method" {
		t.Fatalf("extension name = %q, want exframe.auth.v1.method", got)
	}
	if got := ext.ContainingMessage().FullName(); got != "google.protobuf.MethodOptions" {
		t.Fatalf("extended message = %q, want google.protobuf.MethodOptions", got)
	}
	if got := ext.Message().FullName(); got != "exframe.auth.v1.MethodAuth" {
		t.Fatalf("extension type = %q, want exframe.auth.v1.MethodAuth", got)
	}
}

func TestMethodAuthFields(t *testing.T) {
	t.Parallel()

	fields := (&MethodAuth{}).ProtoReflect().Descriptor().Fields()
	tests := []struct {
		name   string
		number protoreflect.FieldNumber
		kind   protoreflect.Kind
		list   bool
	}{
		{name: "scopes", number: 1, kind: protoreflect.StringKind, list: true},
		{name: "public", number: 2, kind: protoreflect.BoolKind},
	}
	if fields.Len() != len(tests) {
		t.Fatalf("fields = %d, want %d", fields.Len(), len(tests))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := fields.ByName(protoreflect.Name(tt.name))
			if f == nil {
				t.Fatalf("field %s not found", tt.name)
			}
			if f.Number() != tt.number || f.Kind() != tt.kind || f.IsList() != tt.list {
				t.Fatalf("field %s = (%d, %s, list %t), want (%d, %s, list %t)",
					tt.name, f.Number(), f.Kind(), f.IsList(), tt.number, tt.kind, tt.list)
			}
		})
	}
}
