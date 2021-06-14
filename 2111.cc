#include <iostream>
#include <fstream>
#include <set>
#include <unordered_set>
#include <vector>
#include <string.h>

using namespace std;

#define MAX_BUCKETS    201

struct Bucket {
	int len;
	int arr[100];
};

class Solution {
public:
		int IsComposedByShortWords(string &str) {
		if (str.length() <= 0) {
			return 0;
		}

		int iStart = 0;
		int iEnd = str.length();

		const char *pStart = str.c_str();
		int i = 1;

		while (i <= iEnd) {
			string strTmp(pStart, pStart + i);

			if (m_set.count(strTmp) > 0) {

				if (i == iEnd) {
					return 1;
				}

				iStart = i;
				string stTailStr(pStart + iStart, pStart + iEnd);
				int iRet = IsComposedByShortWords(stTailStr);
				if(iRet > 0)
				{
					return iRet+1;
				}
			}
			i++;
		}

		return -1;
	}

	string longestWord(vector<string> &words) {

		bzero(m_buckets, sizeof(m_buckets));

		for (int i = 0; i < words.size(); ++i) {
			m_set.insert(words[i]);

            //构建桶数组
			int iStrLen = words[i].length();
			Bucket *pBucket = &m_buckets[iStrLen];
			string &stCurStr = words[i];

			if (pBucket->len == 0) {
				pBucket->arr[pBucket->len++] = i;
				continue;
			}

			//根据字典序插入排序
			int iInsertIndex = 0;
			for (int k = pBucket->len-1; k >= 0; --k) {
				string &stTmpStr = words[pBucket->arr[k]];
				if (strcmp(stCurStr.c_str(), stTmpStr.c_str()) > 0) {
					pBucket->arr[k + 1] = pBucket->arr[k];
				} else {
					iInsertIndex = k + 1;
					break;
				}
			}
			pBucket->arr[iInsertIndex] = i;
			pBucket->len++;
		}

		for (int i = MAX_BUCKETS; i > 0; --i) {
			Bucket *pBucket = &m_buckets[i];
			if (pBucket->len <= 0) {
				continue;
			}

			for (int k = pBucket->len - 1; k >= 0; --k) {
				string &strTmp = words[pBucket->arr[k]];
				int iComposeCnt = IsComposedByShortWords(strTmp);
				if (iComposeCnt > 1) {
					return strTmp;
				}
			}
		}

		return string("");
	}

	unordered_set<string> m_set;
	Bucket m_buckets[MAX_BUCKETS];
};

int main() {
	fstream file("./block1");
	if (!file.is_open()) {
		cout << "error in open input file!" << endl;
		return -1;
	}

	vector<string> vec;
	string line;
	while (getline(file, line)) {
		vec.push_back(line);
	}

	Solution solu;
	cout <<"result:"<< solu.longestWord(vec) << endl;;

	return 0;
}
