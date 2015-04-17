// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package java

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"v.io/x/ref/lib/vdl/compile"
	"v.io/x/ref/lib/vdl/vdlutil"
)

const serverInterfaceTmpl = header + `
// Source: {{ .Source }}
package {{ .PackagePath }};

{{ .ServerDoc }}
@io.v.v23.vdl.VServer(
    serverWrapper = {{ .ServerWrapperPath }}.class
)
{{ .AccessModifier }} interface {{ .ServiceName }}Server {{ .Extends }} {
{{ range $method := .Methods }}
    {{/* Generate the method signature. */}}
    {{ $method.Doc }}
    {{ $method.AccessModifier }} {{ $method.RetType }} {{ $method.Name }}(final io.v.v23.context.VContext ctx, final io.v.v23.rpc.ServerCall call{{ $method.Args }}) throws io.v.v23.verror.VException;
{{ end }}
}
`

func serverInterfaceOutArg(method *compile.Method, iface *compile.Interface, env *compile.Env) string {
	switch len(method.OutArgs) {
	case 0:
		return "void"
	case 1:
		return javaType(method.OutArgs[0].Type, false, env)
	default:
		return javaPath(path.Join(interfaceFullyQualifiedName(iface)+"Client", method.Name+"Out"))
	}
}

type serverInterfaceMethod struct {
	AccessModifier string
	Args           string
	Doc            string
	Name           string
	RetType        string
}

func processServerInterfaceMethod(method *compile.Method, iface *compile.Interface, env *compile.Env) serverInterfaceMethod {
	args := javaDeclarationArgStr(method.InArgs, env, true)
	if isStreamingMethod(method) {
		args += fmt.Sprintf(", io.v.v23.vdl.Stream<%s, %s> stream", javaType(method.OutStream, true, env), javaType(method.InStream, true, env))
	}
	return serverInterfaceMethod{
		AccessModifier: accessModifierForName(method.Name),
		Args:           args,
		Doc:            method.Doc,
		Name:           vdlutil.FirstRuneToLower(method.Name),
		RetType:        serverInterfaceOutArg(method, iface, env),
	}
}

// genJavaServerInterfaceFile generates the Java interface file for the provided
// interface.
func genJavaServerInterfaceFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	methods := make([]serverInterfaceMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processServerInterfaceMethod(method, iface, env)
	}
	javaServiceName := vdlutil.FirstRuneToUpper(iface.Name)
	data := struct {
		FileDoc           string
		AccessModifier    string
		Extends           string
		Methods           []serverInterfaceMethod
		PackagePath       string
		ServerDoc         string
		ServerVDLPath     string
		ServiceName       string
		ServerWrapperPath string
		Source            string
	}{
		FileDoc:           iface.File.Package.FileDoc,
		AccessModifier:    accessModifierForName(iface.Name),
		Extends:           javaServerExtendsStr(iface.Embeds),
		Methods:           methods,
		PackagePath:       javaPath(javaGenPkgPath(iface.File.Package.GenPath)),
		ServerDoc:         javaDoc(iface.Doc),
		ServiceName:       javaServiceName,
		ServerVDLPath:     path.Join(iface.File.Package.GenPath, iface.Name+"ServerMethods"),
		ServerWrapperPath: javaPath(javaGenPkgPath(path.Join(iface.File.Package.GenPath, javaServiceName+"ServerWrapper"))),
		Source:            iface.File.BaseName,
	}
	var buf bytes.Buffer
	err := parseTmpl("server interface", serverInterfaceTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute struct template: %v", err)
	}
	return JavaFileInfo{
		Name: javaServiceName + "Server.java",
		Data: buf.Bytes(),
	}
}
