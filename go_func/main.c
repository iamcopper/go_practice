#include <stdio.h>

int add(int a, int b);

int main(int argc, char *argv[])
{
	int sum;

	sum = add(1, 2);
	printf("sum = %d\n", sum);

	return 0;
}

int add(int a, int b)
{
	return a + b;
}
