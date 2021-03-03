<!--
---
linkTitle: "Step Composition Contract"
weight: 8
---
-->

# Tekton Step Composition v0.1

This document outlines the step composition model.

## Aims

* Allow Tasks and Steps to be inlined from versioned files in git (or other sources).
* Add local steps before/after/in between the steps from a shared Task
* Override properties on an inherited step such as changing the image, command, args, environment variables, volumes, script.
* Try be fairly DRY yet simple so that its immediately obvious looking at a Tekton YAML what it means
* Inline the override syntax in the controller or a `mutatingwebhookconfigurations`


