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

const clientStubTmpl = header + `
// Source(s):  {{ .Source }}
package {{ .PackagePath }};

/* Client stub for interface: {{ .ServiceName }}Client. */
{{ .AccessModifier }} final class {{ .ServiceName }}ClientStub implements {{ .FullServiceName }}Client {
    private final io.v.v23.rpc.Client client;
    private final java.lang.String vName;

    {{/* Define fields to hold each of the embedded object stubs*/}}
    {{ range $embed := .Embeds }}
    {{/* e.g. private final com.somepackage.gen_impl.ArithStub stubArith; */}}
    private final {{ $embed.StubClassName }} {{ $embed.LocalStubVarName }};
    {{ end }}

    public {{ .ServiceName }}ClientStub(final io.v.v23.rpc.Client client, final java.lang.String vName) {
        this.client = client;
        this.vName = vName;
        {{/* Initialize the embeded stubs */}}
        {{ range $embed := .Embeds }}
        this.{{ $embed.LocalStubVarName }} = new {{ $embed.StubClassName }}(client, vName);
         {{ end }}
    }

    private io.v.v23.rpc.Client getClient(io.v.v23.context.VContext context) {
        return this.client != null ? client : io.v.v23.V.getClient(context);

    }

    // Methods from interface UniversalServiceMethods.
    @Override
    public io.v.v23.vdlroot.signature.Interface getSignature(io.v.v23.context.VContext context) throws io.v.v23.verror.VException {
        return getSignature(context, null);
    }
    @Override
    public io.v.v23.vdlroot.signature.Interface getSignature(io.v.v23.context.VContext context, io.v.v23.Options vOpts) throws io.v.v23.verror.VException {
        // Start the call.
        final io.v.v23.rpc.Client.Call _call = getClient(context).startCall(context, this.vName, "signature", new java.lang.Object[0], new java.lang.reflect.Type[0], vOpts);

        // Finish the call.
        final java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{
            new com.google.common.reflect.TypeToken<io.v.v23.vdlroot.signature.Interface>() {}.getType(),
        };
        final java.lang.Object[] _results = _call.finish(_resultTypes);
        return (io.v.v23.vdlroot.signature.Interface)_results[0];
    }

    // Methods from interface {{ .ServiceName }}Client.
{{/* Iterate over methods defined directly in the body of this service */}}
{{ range $method := .Methods }}
    {{/* The optionless overload simply calls the overload with options */}}
    @Override
    {{ $method.AccessModifier }} {{ $method.RetType }} {{ $method.Name }}(final io.v.v23.context.VContext context{{ $method.DeclarationArgs }}) throws io.v.v23.verror.VException {
        {{if $method.Returns }}return{{ end }} {{ $method.Name }}(context{{ $method.CallingArgsLeadingComma }}, null);
    }
    {{/* The main client stub method body */}}
    @Override
    {{ $method.AccessModifier }} {{ $method.RetType }} {{ $method.Name }}(final io.v.v23.context.VContext context{{ $method.DeclarationArgs }}, io.v.v23.Options vOpts) throws io.v.v23.verror.VException {
        {{/* Start the vanadium call */}}
        // Start the call.
        final java.lang.Object[] _args = new java.lang.Object[]{ {{ $method.CallingArgs }} };
        final java.lang.reflect.Type[] _argTypes = new java.lang.reflect.Type[]{ {{ $method.CallingArgTypes }} };
        final io.v.v23.rpc.Client.Call _call = getClient(context).startCall(context, this.vName, "{{ $method.Name }}", _args, _argTypes, vOpts);

        // Finish the call.
        {{/* Now handle returning from the function. */}}
        {{ if $method.NotStreaming }}

        {{ if $method.IsVoid }}
        final java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{};
        _call.finish(_resultTypes);
        {{ else }} {{/* else $method.IsVoid */}}
        final java.lang.reflect.Type[] _resultTypes = new java.lang.reflect.Type[]{
            {{ range $outArg := $method.OutArgs }}
            new com.google.common.reflect.TypeToken<{{ $outArg.Type }}>() {}.getType(),
            {{ end }}
        };
        final java.lang.Object[] _results = _call.finish(_resultTypes);
        {{ if $method.MultipleReturn }}
        final {{ $method.DeclaredObjectRetType }} _ret = new {{ $method.DeclaredObjectRetType }}();
            {{ range $i, $outArg := $method.OutArgs }}
        _ret.{{ $outArg.FieldName }} = ({{ $outArg.Type }})_results[{{ $i }}];
            {{ end }} {{/* end range over outargs */}}
        return _ret;
        {{ else }} {{/* end if $method.MultipleReturn */}}
        return ({{ $method.DeclaredObjectRetType }})_results[0];
        {{ end }} {{/* end if $method.MultipleReturn */}}

        {{ end }} {{/* end if $method.IsVoid */}}

        {{else }} {{/* else $method.NotStreaming */}}
        return new io.v.v23.vdl.ClientStream<{{ $method.SendType }}, {{ $method.RecvType }}, {{ $method.DeclaredObjectRetType }}>() {
            @Override
            public void send(final {{ $method.SendType }} item) throws io.v.v23.verror.VException {
                final java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.SendType }}>() {}.getType();
                _call.send(item, type);
            }
            @Override
            public {{ $method.RecvType }} recv() throws java.io.EOFException, io.v.v23.verror.VException {
                final java.lang.reflect.Type type = new com.google.common.reflect.TypeToken<{{ $method.RecvType }}>() {}.getType();
                final java.lang.Object result = _call.recv(type);
                try {
                    return ({{ $method.RecvType }})result;
                } catch (java.lang.ClassCastException e) {
                    throw new io.v.v23.verror.VException("Unexpected result type: " + result.getClass().getCanonicalName());
                }
            }
            @Override
            public {{ $method.DeclaredObjectRetType }} finish() throws io.v.v23.verror.VException {
                {{ if $method.IsVoid }}
                final java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{};
                _call.finish(resultTypes);
                return null;
                {{ else }} {{/* else $method.IsVoid */}}
                final java.lang.reflect.Type[] resultTypes = new java.lang.reflect.Type[]{
                    new com.google.common.reflect.TypeToken<{{ $method.DeclaredObjectRetType }}>() {}.getType()
                };
                return ({{ $method.DeclaredObjectRetType }})_call.finish(resultTypes)[0];
                {{ end }} {{/* end if $method.IsVoid */}}
            }
        };
        {{ end }}{{/* end if $method.NotStreaming */}}
    }
{{ end }}{{/* end range over methods */}}

{{/* Iterate over methods from embeded services and generate code to delegate the work */}}
{{ range $eMethod := .EmbedMethods }}
    @Override
    {{ $eMethod.AccessModifier }} {{ $eMethod.RetType }} {{ $eMethod.Name }}(final io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}) throws io.v.v23.verror.VException {
        {{/* e.g. return this.stubArith.cosine(context, [args]) */}}
        {{ if $eMethod.Returns }}return{{ end }} this.{{ $eMethod.LocalStubVarName }}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }});
    }
    @Override
    {{ $eMethod.AccessModifier }} {{ $eMethod.RetType }} {{ $eMethod.Name }}(final io.v.v23.context.VContext context{{ $eMethod.DeclarationArgs }}, io.v.v23.Options vOpts) throws io.v.v23.verror.VException {
        {{/* e.g. return this.stubArith.cosine(context, [args], options) */}}
        {{ if $eMethod.Returns }}return{{ end }}  this.{{ $eMethod.LocalStubVarName }}.{{ $eMethod.Name }}(context{{ $eMethod.CallingArgsLeadingComma }}, vOpts);
    }
{{ end }}

}
`

type clientStubMethodOutArg struct {
	FieldName string
	Type      string
}

type clientStubMethod struct {
	AccessModifier          string
	CallingArgs             string
	CallingArgTypes         string
	CallingArgsLeadingComma string
	DeclarationArgs         string
	DeclaredObjectRetType   string
	IsVoid                  bool
	MultipleReturn          bool
	Name                    string
	NotStreaming            bool
	OutArgs                 []clientStubMethodOutArg
	RecvType                string
	RetType                 string
	Returns                 bool
	SendType                string
	ServiceName             string
}

type clientStubEmbedMethod struct {
	AccessModifier          string
	CallingArgsLeadingComma string
	DeclarationArgs         string
	LocalStubVarName        string
	Name                    string
	RetType                 string
	Returns                 bool
}

type clientStubEmbed struct {
	StubClassName    string
	LocalStubVarName string
}

func processClientStubMethod(iface *compile.Interface, method *compile.Method, env *compile.Env) clientStubMethod {
	outArgs := make([]clientStubMethodOutArg, len(method.OutArgs))
	for i := 0; i < len(method.OutArgs); i++ {
		if method.OutArgs[i].Name != "" {
			outArgs[i].FieldName = vdlutil.FirstRuneToLower(method.OutArgs[i].Name)
		} else {
			outArgs[i].FieldName = fmt.Sprintf("ret%d", i+1)
		}
		outArgs[i].Type = javaType(method.OutArgs[i].Type, true, env)
	}
	return clientStubMethod{
		AccessModifier:          accessModifierForName(method.Name),
		CallingArgs:             javaCallingArgStr(method.InArgs, false),
		CallingArgTypes:         javaCallingArgTypeStr(method.InArgs, env),
		CallingArgsLeadingComma: javaCallingArgStr(method.InArgs, true),
		DeclarationArgs:         javaDeclarationArgStr(method.InArgs, env, true),
		DeclaredObjectRetType:   clientInterfaceNonStreamingOutArg(iface, method, true, env),
		IsVoid:                  len(method.OutArgs) < 1,
		MultipleReturn:          len(method.OutArgs) > 1,
		Name:                    vdlutil.FirstRuneToLower(method.Name),
		NotStreaming:            !isStreamingMethod(method),
		OutArgs:                 outArgs,
		RecvType:                javaType(method.OutStream, true, env),
		RetType:                 clientInterfaceOutArg(iface, method, false, env),
		Returns:                 len(method.OutArgs) >= 1 || isStreamingMethod(method),
		SendType:                javaType(method.InStream, true, env),
		ServiceName:             vdlutil.FirstRuneToUpper(iface.Name),
	}
}

func processClientStubEmbedMethod(iface *compile.Interface, embedMethod *compile.Method, env *compile.Env) clientStubEmbedMethod {
	return clientStubEmbedMethod{
		AccessModifier:          accessModifierForName(embedMethod.Name),
		CallingArgsLeadingComma: javaCallingArgStr(embedMethod.InArgs, true),
		DeclarationArgs:         javaDeclarationArgStr(embedMethod.InArgs, env, true),
		LocalStubVarName:        vdlutil.FirstRuneToLower(iface.Name) + "ClientStub",
		Name:                    vdlutil.FirstRuneToLower(embedMethod.Name),
		RetType:                 clientInterfaceOutArg(iface, embedMethod, false, env),
		Returns:                 len(embedMethod.OutArgs) >= 1 || isStreamingMethod(embedMethod),
	}
}

// genJavaClientStubFile generates a client stub for the specified interface.
func genJavaClientStubFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	embeds := []clientStubEmbed{}
	for _, embed := range allEmbeddedIfaces(iface) {
		embeds = append(embeds, clientStubEmbed{
			LocalStubVarName: vdlutil.FirstRuneToLower(embed.Name) + "ClientStub",
			StubClassName:    javaPath(javaGenPkgPath(path.Join(embed.File.Package.GenPath, vdlutil.FirstRuneToUpper(embed.Name)+"ClientStub"))),
		})
	}
	embedMethods := []clientStubEmbedMethod{}
	for _, embedMao := range dedupedEmbeddedMethodAndOrigins(iface) {
		embedMethods = append(embedMethods, processClientStubEmbedMethod(embedMao.Origin, embedMao.Method, env))
	}
	methods := make([]clientStubMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processClientStubMethod(iface, method, env)
	}
	javaServiceName := vdlutil.FirstRuneToUpper(iface.Name)
	data := struct {
		FileDoc          string
		AccessModifier   string
		EmbedMethods     []clientStubEmbedMethod
		Embeds           []clientStubEmbed
		FullServiceName  string
		Methods          []clientStubMethod
		PackagePath      string
		ServiceName      string
		Source           string
		VDLIfacePathName string
	}{
		FileDoc:          iface.File.Package.FileDoc,
		AccessModifier:   accessModifierForName(iface.Name),
		EmbedMethods:     embedMethods,
		Embeds:           embeds,
		FullServiceName:  javaPath(interfaceFullyQualifiedName(iface)),
		Methods:          methods,
		PackagePath:      javaPath(javaGenPkgPath(iface.File.Package.GenPath)),
		ServiceName:      javaServiceName,
		Source:           iface.File.BaseName,
		VDLIfacePathName: path.Join(iface.File.Package.GenPath, iface.Name+"ClientMethods"),
	}
	var buf bytes.Buffer
	err := parseTmpl("client stub", clientStubTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute client stub template: %v", err)
	}
	return JavaFileInfo{
		Name: javaServiceName + "ClientStub.java",
		Data: buf.Bytes(),
	}
}
