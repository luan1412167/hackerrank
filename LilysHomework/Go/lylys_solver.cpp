#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);
vector<string> split(const string &);

/*
 * Complete the 'lilysHomework' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

int lilysHomework(vector<int> arr) {
    vector<int> arr_org(arr);
    vector<int> arr_inc(arr);
    vector<int> arr_dec(arr);
    vector<int> sorted_arr_inc(arr);
    sort(sorted_arr_inc.begin(), sorted_arr_inc.end());
    vector<int> sorted_arr_dec(sorted_arr_inc);
    reverse(sorted_arr_dec.begin(), sorted_arr_dec.end());
        
    unordered_map<int, int> mapArr;
    for (int i=0; i < arr_org.size(); i++)
    {
        mapArr[arr_org[i]] = i;
    }
    unordered_map<int ,int> mapArrDec = mapArr;
    int ret_inc = 0;
    int ret_dec = 0;
    for (int i=0; i < arr_inc.size(); i++)
    {
        if(sorted_arr_inc[i] != arr_inc[i]){
            ret_inc++;
            int desire_val = sorted_arr_inc[i];
            int swap_val = arr_inc[i];
            int desire_val_idx = mapArr[desire_val];
            int swap_val_idx = i;
            arr_inc[swap_val_idx] = desire_val;
            arr_inc[desire_val_idx] = swap_val;
            mapArr[desire_val] = swap_val_idx;
            mapArr[swap_val] = desire_val_idx;
        }
    }
    
    for (int i=0; i < arr_dec.size(); i++)
    {
        if(sorted_arr_dec[i] != arr_dec[i]){
            ret_dec++;
            int desire_val = sorted_arr_dec[i];
            int swap_val = arr_dec[i];
            int desire_val_idx = mapArrDec[desire_val];
            int swap_val_idx = i;
            arr_dec[swap_val_idx] = desire_val;
            arr_dec[desire_val_idx] = swap_val;
            mapArrDec[desire_val] = swap_val_idx;
            mapArrDec[swap_val] = desire_val_idx;
        }
    }
    
    return min(ret_inc, ret_dec);
    
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string n_temp;
    getline(cin, n_temp);

    int n = stoi(ltrim(rtrim(n_temp)));

    string arr_temp_temp;
    getline(cin, arr_temp_temp);

    vector<string> arr_temp = split(rtrim(arr_temp_temp));

    vector<int> arr(n);

    for (int i = 0; i < n; i++) {
        int arr_item = stoi(arr_temp[i]);

        arr[i] = arr_item;
    }

    int result = lilysHomework(arr);

    fout << result << "\n";

    fout.close();

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}

vector<string> split(const string &str) {
    vector<string> tokens;

    string::size_type start = 0;
    string::size_type end = 0;

    while ((end = str.find(" ", start)) != string::npos) {
        tokens.push_back(str.substr(start, end - start));

        start = end + 1;
    }

    tokens.push_back(str.substr(start));

    return tokens;
}