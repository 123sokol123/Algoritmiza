#include<iostream>

using namespace std;

int main() {
	int n;
	cin >> n;

	int maxLength = 2, c = 1, direction = 1, currentLength = 1;
	bool shouldExit = false;

	while (true) {
		for (int j = 0; j < currentLength; j++) {
			cout << c << " ";

			if (c == n) {
				shouldExit = true;
				break;
			}
			c++;
		}
		cout << "\n";
		if (shouldExit) break;

		currentLength += direction;
		if (currentLength == maxLength) {
			direction = -1;
			maxLength++;
		}

		if (currentLength == 1) {
			direction = +1;
		}
	}

	return 0;
}