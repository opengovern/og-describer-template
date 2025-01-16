package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/opengovern/og-describer-github/global"
	"os"
	"sort"
	"strings"
	"text/template"
)

// Define the ResourceType struct with Labels and Annotations
type ResourceType struct {
	ResourceName      string
	Tags              map[string][]string
	TagsString        string `json:"-"`
	ListDescriber     string
	GetDescriber      string
	SteampipeTable    string
	Model             string
	Annotations       map[string]string
	Labels            map[string]string
	AnnotationsString string `json:"-"`
	LabelsString      string `json:"-"`
}

var (
	resourceTypesFile = flag.String("resourceTypesFile", "", "Location of the resource types json file file")
	output            = flag.String("output", "", "Path to the output file for resource types")
	resourceTypesList = flag.String("resource-types-list", "", "Path to the output file for index map")
)

func main() {
	flag.Parse()

	var resourceTypes []ResourceType

	if resourceTypesFile == nil || len(*resourceTypesFile) == 0 {
		rt := "global/maps/resource-types.json"
		resourceTypesFile = &rt
	}

	resourceTypesFileJson, err := os.ReadFile(*resourceTypesFile)
	if err != nil {
		panic(err)
	}
	// Parse the embedded JSON into resourceTypes slice
	if err := json.Unmarshal(resourceTypesFileJson, &resourceTypes); err != nil {
		panic(err)
	}

	// Define the template with Labels and Annotations included
	tmpl, err := template.New("").Parse(fmt.Sprintf(`
	"{{ .ResourceName }}": {
		IntegrationType:      IntegrationName,
		ResourceName:         "{{ .ResourceName }}",
		Tags:                 {{ .TagsString }},
		Labels:               {{ .LabelsString }},
		Annotations:          {{ .AnnotationsString }},
		ListDescriber:        provider.{{ .ListDescriber }},
		GetDescriber:         {{ if .GetDescriber }}provider.{{ .GetDescriber }}{{ else }}nil{{ end }},
	},
`))
	if err != nil {
		panic(err)
	}

	// Set default output paths if not provided
	if output == nil || len(*output) == 0 {
		v := "global/maps/provider_resource_types.go"
		output = &v
	}

	if resourceTypesList == nil || len(*resourceTypesList) == 0 {
		v := "resource_types_list.go"
		resourceTypesList = &v
	}

	// Initialize a strings.Builder to construct the output file content
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf(`package maps
import (
	"%[1]s/discovery/describers"
	"%[1]s/discovery/provider"
	model "github.com/opengovern/og-describer-%[2]s/discovery/pkg/models"
	 "github.com/opengovern/og-util/pkg/integration"
)

const (
	IntegrationName      = integration.Type("%[3]s")          
)
var ResourceTypes = map[string]model.ResourceType{
`, global.OGPluginRepoURL, global.IntegrationTypeLower,global.IntegrationName))

	// Iterate over each resource type to build its string representations
	for _, resourceType := range resourceTypes {
		var arr []string

		// Build TagsString
		tagsStringBuilder := strings.Builder{}
		tagsStringBuilder.WriteString("map[string][]string{\n")

		var tagsLines []string
		for k, v := range resourceType.Tags {
			arr = []string{}
			for _, t := range v {
				arr = append(arr, "\""+t+"\"")
			}
			tagsLines = append(tagsLines, fmt.Sprintf("            \"%s\": {%s},\n", k, strings.Join(arr, ",")))
		}

		sort.Strings(tagsLines) // Sort for consistency
		for _, l := range tagsLines {
			tagsStringBuilder.WriteString(l)
		}

		tagsStringBuilder.WriteString("        }")
		resourceType.TagsString = tagsStringBuilder.String()

		// Build LabelsString
		labelsStringBuilder := strings.Builder{}
		labelsStringBuilder.WriteString("map[string]string{\n")

		var labelsLines []string
		for k, v := range resourceType.Labels {
			// Escape quotes in keys and values
			escapedKey := escapeString(k)
			escapedValue := escapeString(v)
			labelsLines = append(labelsLines, fmt.Sprintf("            \"%s\": \"%s\",\n", escapedKey, escapedValue))
		}

		sort.Strings(labelsLines) // Sort for consistency
		for _, l := range labelsLines {
			labelsStringBuilder.WriteString(l)
		}

		labelsStringBuilder.WriteString("        }")
		resourceType.LabelsString = labelsStringBuilder.String()

		// Build AnnotationsString
		annotationsStringBuilder := strings.Builder{}
		annotationsStringBuilder.WriteString("map[string]string{\n")

		var annotationsLines []string
		for k, v := range resourceType.Annotations {
			// Escape quotes in keys and values
			escapedKey := escapeString(k)
			escapedValue := escapeString(v)
			annotationsLines = append(annotationsLines, fmt.Sprintf("            \"%s\": \"%s\",\n", escapedKey, escapedValue))
		}

		sort.Strings(annotationsLines) // Sort for consistency
		for _, l := range annotationsLines {
			annotationsStringBuilder.WriteString(l)
		}

		annotationsStringBuilder.WriteString("        }")
		resourceType.AnnotationsString = annotationsStringBuilder.String()

		// Execute the template with the current resourceType
		err = tmpl.Execute(b, resourceType)
		if err != nil {
			panic(err)
		}
	}
	b.WriteString("}\n")

	// Write the generated content to the output file
	err = os.WriteFile(*output, []byte(b.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Generate the index map file as before
	b = &strings.Builder{}
	b.WriteString(fmt.Sprintf(`package configs

var ResourceTypesList = []string{
`))
	for _, resourceType := range resourceTypes {
		b.WriteString(fmt.Sprintf("  \"%s\",\n", resourceType.ResourceName))
	}
	b.WriteString(fmt.Sprintf(`}`))

	// Write the index map to the specified file
	err = os.WriteFile(*resourceTypesList, []byte(b.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// escapeString ensures that any quotes in the strings are properly escaped
func escapeString(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
