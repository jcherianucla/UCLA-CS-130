#include<iostream>

using namespace std;

int sum(int x, int y) {
	return x + y;
}

int diff(int x, int y) {
	return x - y;
}

int answerToLife() {
	return 42;
}

int main(int argc, char **argv) {
	cout << answerToLife() << endl;
	return 0;
}
