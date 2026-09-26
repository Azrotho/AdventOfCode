#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#define ROW_NUMBER 135
#define LINE_NUMBER 135

using namespace std;


void corners(const vector<vector<char>>& mat, int rows, int cols) {
    if (rows == 0 || cols == 0) return;
    
    cout << "topleft: " << (char)mat[0][0] << endl;
    cout << "topright: " << (char)mat[0][cols-1] << endl;
    cout << "bottomleft: " << (char)mat[rows-1][0] << endl;
    cout << "bottomright: " << (char)mat[rows-1][cols-1] << endl;
}

int main() {
    ifstream file("input.txt");

    string line;


    std::vector<std::vector<char> > mat(ROW_NUMBER);
    int nbline = 0;
    int total = 0;
    int previoustotal = 1;
    for (int i = 0; i < ROW_NUMBER; i++) {
        mat[i].resize(LINE_NUMBER); // this will allow you to now just use [][] to access stuff
    }

    if (file.is_open()) {
        while (getline(file, line)) {
            for(int i = 0; i < line.length(); i++) {
                mat[nbline][i] = line[i];
            }
            nbline++;
        }
        
        // traitement en boucle dans la matrice
        while(total - previoustotal != 0) {
            previoustotal = total;
            for(int i = 0; i < ROW_NUMBER; i++) {
                for(int j = 0; j < LINE_NUMBER; j++) {
                    if(mat[i][j] == '@') {
                        int voisins = 0;
                        // Comptons les voisins @
                        for(int i2 = 0; i2 < 3; i2++) {
                            for(int j2 = 0; j2 < 3; j2++) {
                                int ni = i + i2 - 1;
                                int nj = j + j2 - 1;
                                if (i2 == 1 && j2 == 1) continue;
                                    if (ni >= 0 && ni < ROW_NUMBER && nj >= 0 && nj < LINE_NUMBER) {
                                        if (mat[ni][nj] == '@') {
                                            voisins++;
                                        }
                                    }
                                }
                            }
                            if(voisins < 4) {
                                mat[i][j] = '.';
                                total++;
                            }
                        }
                    }
                }
            }
        cout << total << endl;
        file.close();
    }
}