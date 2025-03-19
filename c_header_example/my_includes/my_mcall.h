#ifndef MY_MCALL_H
#define MY_MCALL_H

#define MAX_LENGTH 128

typedef enum mcops_e
{
    MCRM_STATVV = 0,
} mcops_e_t;

typedef struct mcall{
    char my_name[MAX_LENGTH];
    int my_age;
    int other_stuff;
    int other_stuff_b;
} my_mcall_t;

#endif // MY_MCALL_H