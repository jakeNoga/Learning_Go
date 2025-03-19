#include "c_serv_calls.h"
#include <stddef.h> // For NULL
#include <stdio.h>  // For stderr and fprintf
#include <stdlib.h> // For malloc and free
#include <string.h> // For strncpy
#include <unistd.h> // For close

// Function to parse input
int parse(int argc, char *argv[], myStruct_t *thisStruct) {
    char *ptr;
    if (argc < 4) {
        return Err;
    }

    // Take arguments in from the command line in the form of -a=<int>, -b=<int>, -s=<string>
    for (int i = 1; i < argc; i++) {
        if (strstr(argv[i], "-a=") != NULL) {
            // split string after = sign
            ptr = strchr(argv[i], '=');
            ptr++;
            // TODO: This is unsafe should use strtol but being lazy
            thisStruct->a = atoi(ptr);
        }
        if (strstr(argv[i], "-b=") != NULL) {
            // split string after = sign
            ptr = strchr(argv[i], '=');
            ptr++;
            // TODO: this is unsafe should use strtol but being lazy
            thisStruct->b = atoi(ptr);
        }
        if (strstr(argv[i], "-s") != NULL) {
            // split string after = sign
            ptr = strchr(argv[i], '=');
            ptr++;

            strncpy(thisStruct->myString, ptr, sizeof(thisStruct->myString));
        }
    }

    return Ok;
}

int main (int argc, char *argv[]) {
    myStruct_t *thisStruct = malloc(sizeof(myStruct_t));

    if (thisStruct == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return -1;
    }
    // set base values of thisStruct
    thisStruct->a = 1;
    thisStruct->b = 2;
    snprintf(thisStruct->myString, sizeof(thisStruct->myString), "Hello World");


    parse(argc, argv, thisStruct);
    callServer(&thisStruct);

    free(thisStruct);
    return 0;
}