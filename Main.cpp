//
// Created by Administrator on 2021/3/27.
//

#include "common/BloomFilter.h"

#include <iostream>

using namespace std;

int main()
{
    unsigned char *pData = new unsigned char[2048];
    
    BloomFilter stFilter;
    stFilter.Init(pData,1024 * WORLD_BLOM_FILTER_HASH_CNT);
    
    uint32_t array[5] = {111111,222222,123456,654321,654};
    for (int i = 0; i < 5; ++i)
    {
        stFilter.Insert(array[i]);
    }
    
    for (int i = 0; i < 5; ++i)
    {
        bool bRet = stFilter.Contains(array[i]);
        cout << array[i] << ":" << bRet << endl;
    }
    
    return 0;
}