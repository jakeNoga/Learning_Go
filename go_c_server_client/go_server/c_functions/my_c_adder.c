#include "myTypes.h"
#include <stdio.h>

int mcall_adder(myStruct_t **thisStruct) {
    myStruct_t* this_stuct = *thisStruct;
    int result = this_stuct->a + this_stuct->b;
    printf("The message was: %s\n", this_stuct->myString);
    printf("The result is: %d\n", result);
    return result;
}