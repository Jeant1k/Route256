#include <iostream>
#include <map>
#include <vector>

using namespace std;

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        int numPlayers;
        cin >> numPlayers;

        int time;
        multimap<int, int> times;
        for (int j = 0; j < numPlayers; ++j) {
            cin >> time;
            times.insert(make_pair(time, j));
        }

        int place = 0, prevTime = -1, numSamePlaces = 1;
        vector<int> places(numPlayers);
        for (auto& [time, idx] : times) {
            if (time - prevTime > 1) {
                place += numSamePlaces;
                numSamePlaces = 1;
            } else ++numSamePlaces;
            places[idx] = place;
            prevTime = time;
        }

        for (int place : places)
            cout << place << ' ';
        cout << endl;

    }

    
    return 0;

}
