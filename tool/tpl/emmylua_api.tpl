{{range $i, $class := .Classes}}
---@class {{$class.Name}} {{range $fieldIndex, $field := $class.Fields}}
---@field public {{$field.Name}} {{GetClassFieldType $field}} {{if hasClassFieldComment $field}} @{{GetClassFieldComment $field}} {{end}}{{end}}
local {{$class.Name}} = {}
{{end}}