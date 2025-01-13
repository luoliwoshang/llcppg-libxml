#include <libxml/parser.h>
#include <libxml/tree.h>
#include <stdio.h>
#include <string.h> // Add this line for strlen()

int main() {
    // XML string to parse
    const char *xmlStr = "<?xml version='1.0'?>"
                         "<root>"
                         "    <person>"
                         "        <name>Alice</name>"
                         "        <age>25</age>"
                         "    </person>"
                         "</root>";

    // Init libxml2
    xmlInitParser();

    // Parse XML string
    xmlDocPtr doc = xmlReadMemory(xmlStr, strlen(xmlStr), NULL, NULL, 0);
    if (doc == NULL) {
        printf("Failed to parse XML\n");
        return 1;
    }

    // Get root element
    xmlNodePtr root = xmlDocGetRootElement(doc);

    // Print root element name
    printf("Root element: %s\n", root->name);

    // Clean up
    xmlFreeDoc(doc);
    xmlCleanupParser();

    return 0;
}