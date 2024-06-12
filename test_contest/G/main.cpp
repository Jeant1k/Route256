#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {

    int numTests;
    cin >> numTests;

    for (int i = 0; i < numTests; ++i) {

        int n, m;
        cin >> n >> m;

        vector<pair<int, int>> appointments(m);
        for (int j = 0; j < m; ++j) {
            cin >> appointments[j].first;
            appointments[j].second = j;
        }

        stable_sort(appointments.begin(), appointments.end());

        bool possible = true;
        int prevAppointment = 0;
        vector<char> appointmentShifts(n);
        for (auto& [appointment, idx] : appointments) {

            if (appointment - 1 > prevAppointment) {
                appointmentShifts[idx] = '-';
                prevAppointment = appointment - 1;
            } else if (appointment > prevAppointment) {
                appointmentShifts[idx] = '0';
                prevAppointment = appointment;
            } else if (appointment + 1 > prevAppointment && appointment + 1 <= n) {
                appointmentShifts[idx] = '+';
                prevAppointment = appointment + 1;
            } else {
                cout << 'x' << endl;
                possible = false;
                break;
            }

        }

        if (possible) {
            for (char shift : appointmentShifts)
                cout << shift;
            cout << endl;
        }

    }

    return 0;

}
