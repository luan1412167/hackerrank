#include <bits/stdc++.h>

using namespace std;

// Complete the substrCount function below.
long substrCount(int n, string s) {
    long int count = n;
    for (int i=0; i<n; i++)
    {
        // count special string with 1 char
        int j = 1;
        while(i+j<n && s[i]==s[i+j])
        {
            count++;
            j++;
        }
        
        // count special with middle one
        
        int k=1;
        
        while (i-k>=0 && i+k<n && s[i]!=s[i+k] && s[i-k] == s[i+k] && s[i+1] == s[i+k]) {
            count++;
            k++;
        }
    }
    return count;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string s;
    getline(cin, s);

    long result = substrCount(n, s);

    fout << result << "\n";

    fout.close();

    return 0;
}
