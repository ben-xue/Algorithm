#ifndef COMM_CWORLDBLOOMFILTER_H
#define COMM_CWORLDBLOOMFILTER_H

#include <inttypes.h>

// pData size calc: ElementCnt / 8 + 1
static const int WORLD_BLOM_FILTER_HASH_CNT = 2;

class BloomFilter
{
public:

    BloomFilter();

    ~BloomFilter() = default;

    bool Init(uint8_t * pData, uint32_t iSize);

    bool IsInit() const;

    void Insert(const uint32_t & data);

    void Insert(const char* data, const uint32_t & length);

    bool Contains(const uint32_t & data) const;

protected:

    bool Contains(const unsigned char* key_begin, const uint32_t length) const;

    void Insert(const unsigned char* key_begin, const uint32_t & length);

    void SetBit(uint32_t iPos);

    bool GetBit(uint32_t iPos) const;

private:
    BloomFilter(const BloomFilter& hfs) = delete;
    BloomFilter& operator=(const BloomFilter& hfs) = delete;

    uint8_t * m_pData;
    uint32_t m_DataSize;
    uint32_t m_InsertedElememtCnt;
};


#endif //COMM_CWORLDBLOOMFILTER_H
