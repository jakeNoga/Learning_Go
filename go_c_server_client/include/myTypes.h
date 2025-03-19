#ifndef _MYTYPES_H_
#define _MYTYPES_H_
#define BUFFER_SIZE 256

typedef struct {
    int a;
    int b;
    char myString[BUFFER_SIZE];
} myStruct_t;

typedef enum {
    Ok = 0,
    Err = 1,
} error_t;

#endif // _MYTYPES_H_