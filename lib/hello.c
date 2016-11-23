#include <stdio.h>
#include "hello.h"

void hello(void)
{
    puts("Hello from dynamic library");
}

int sum(int x, int y)
{
    return x + y;
}