package pepperlint

import (
	"fmt"
	"go/ast"
)

type testExcludeNameTypeSpecRule struct {
	Name string
}

func (v testExcludeNameTypeSpecRule) ValidateTypeSpec(spec *ast.TypeSpec) error {
	if spec.Name.Name == v.Name {
		return fmt.Errorf("%s", v.Name)
	}

	return nil
}

type testExcludeField struct {
	Name string
}

func (v testExcludeField) ValidateStructType(s *ast.StructType) error {
	for _, field := range s.Fields.List {
		for _, name := range field.Names {
			if name.Name == v.Name {
				return fmt.Errorf("%s", name)
			}
		}
	}

	return nil
}

type testExcludeMethod struct {
	StructName string
	Name       string
	helper     Helper
}

func (v testExcludeMethod) ValidateFuncDecl(fnDecl *ast.FuncDecl) error {
	if !IsMethod(fnDecl) {
		return nil
	}

	found := false
	for _, object := range fnDecl.Recv.List {
		if ok := v.helper.IsStruct(object.Type); !ok {
			continue
		}

		name := v.helper.GetStructName(object.Type)
		if name == v.StructName {
			found = true
		}
	}

	if !found {
		return nil
	}

	if fnDecl.Name.Name == v.Name {
		return fmt.Errorf("%s", v.Name)
	}

	return nil
}
