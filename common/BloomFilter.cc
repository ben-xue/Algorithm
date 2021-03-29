
#include "BloomFilter.h"
#include "MurmurHash3.h"


BloomFilter::BloomFilter() :
        m_pData(0),
        m_DataSize(0),
        m_InsertedElememtCnt(0)
{

}

bool BloomFilter::Init(uint8_t * pData, uint32_t iSize)
{
    m_pData = pData;
    m_DataSize = iSize;

    return IsInit();
}

bool BloomFilter::IsInit() const
{
    if (m_pData && m_DataSize)
    {
        return true;
    }

    return false;
}

void BloomFilter::Insert(const uint32_t &data)
{
    Insert(reinterpret_cast<const unsigned char*>(&data),sizeof(data));
}

void BloomFilter::Insert(const char *data, const uint32_t &length)
{
    Insert(reinterpret_cast<const unsigned char*>(data),length);
}

bool BloomFilter::Contains(const uint32_t &data) const
{
    return Contains(reinterpret_cast<const unsigned char*>(&data),static_cast<uint32_t>(sizeof(data)));
}

bool BloomFilter::Contains(const unsigned char *key_begin, const uint32_t length) const
{
    uint32_t seed = 0;
    for (int i = 0; i < WORLD_BLOM_FILTER_HASH_CNT; i++)
    {
        uint32_t pos = 0;
        MurmurHash3_x86_32(key_begin, length, seed, &pos);
        seed = pos;
        pos %= m_DataSize;

        if (!GetBit(pos))
        {
            return false;
        }
    }

    return true;
}

void BloomFilter::Insert(const unsigned char *key_begin, const uint32_t &length)
{
    uint32_t seed = 0;
    for(int i=0; i< WORLD_BLOM_FILTER_HASH_CNT; i++)
    {
        uint32_t pos = 0;
        MurmurHash3_x86_32(key_begin, length, seed, &pos);
        seed = pos;
        pos %= m_DataSize;
        SetBit(pos);
    }

    ++m_InsertedElememtCnt;
}

void BloomFilter::SetBit(uint32_t iPos)
{
    uint32_t iByteIndex = (uint32_t)(iPos >> 3);
    m_pData[iByteIndex] |= (1 << (iPos & 7));
}

bool BloomFilter::GetBit(uint32_t iPos) const
{
    uint32_t iByteIndex = (uint32_t)(iPos >> 3);
    return (m_pData[iByteIndex] & (1 << (iPos & 7)));
}

