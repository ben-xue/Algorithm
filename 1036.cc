
/*
在一个 10^6 x 10^6 的网格中，每个网格上方格的坐标为 (x, y) 。

现在从源方格 source = [sx, sy] 开始出发，意图赶往目标方格 target = [tx, ty] 。

数组 blocked 是封锁的方格列表，其中每个 blocked[i] = [xi, yi] 表示坐标为 (xi, yi) 的方格是禁止通行的。

每次移动，都可以走到网格中在四个方向上相邻的方格，只要该方格 不 在给出的封锁列表 blocked 上。同时，不允许走出网格。

只有在可以通过一系列的移动从源方格 source 到达目标方格 target 时才返回 true。否则，返回 false。

*/

#include <iostream>
#include <vector>
#include <map>
#include <set>

using namespace std;

template <typename KeyT, typename ValueT>
class CIndexedBinaryHeap
{
public:

	CIndexedBinaryHeap<KeyT, ValueT>()
	{

	}
  
	int Push(ValueT& elem)
	{
		m_value.push_back(elem);
		m_keyIndex[elem.key] = m_value.size() - 1;
		return FixUp(elem.key, m_value.size() - 1);
	}

	bool Empty()
	{
		return m_value.size() == 0;
	}

	bool Exist(const KeyT& key)
	{
		return m_keyIndex.find(key) != m_keyIndex.end();
	}

	ValueT& Get(const KeyT& key)
	{
		if (Exist(key))
		{
			int idx = m_keyIndex[key];
			return m_value[idx];
		}
		
		throw "key not found exception";
	}

	ValueT Pop()
	{
		if (Empty())
		{
			throw "empty heap can not pop";
		}

		ValueT v = m_value[0];
		swap(m_value[0], m_value[m_value.size() - 1]);
		m_value.pop_back();
		m_keyIndex.erase(v.key);

		if (!Empty())
		{
			m_keyIndex[m_value[0].key] = 0;	//value下标变化，索引到跟着调整
			FixDown(m_value[0].key, 0);
		}

		return v;
	}
	
	int Update(const KeyT& key)
	{
		if (m_keyIndex.find(key) == m_keyIndex.end())
		{
			return -1;
		}
		
		int idx = m_keyIndex[key];
		
		int f = (idx + 1) / 2 - 1;
        // 有父节点并且父节点大
		if (f >= 0 && m_value[idx] < m_value[f])	
		{
			return FixUp(key, idx);
		}

        return FixDown(key, idx);
	}

private:
	vector<ValueT> m_value;
	map<KeyT, int> m_keyIndex;

	int FixDown(const KeyT& key, int idx)
	{
		int f = idx;
		int k = 0;
		while ((k = (f * 2) + 1) < m_value.size())
		{ 
			if (k < m_value.size() - 1)	//存在右子节点
			{
				if (m_value[k + 1] < m_value[k])
				{
					k++;
				}
			}
            
			if (m_value[f] <= m_value[k]) 				//father已经最小了
			{
				break;
			}

			KeyT fk = m_value[f].key;
			KeyT kk = m_value[k].key;

			swap(m_value[f], m_value[k]);
			swap(m_keyIndex[fk], m_keyIndex[kk]);
			f = k;
		}
		return k;
	}

	int FixUp(const KeyT& key, int idx)
	{
		int c = idx;
		int f = 0;
		while ((f =  (c + 1) / 2 - 1 )>= 0)
		{
			if (m_value[f] <= m_value[c])
			{
				break;
			}

			KeyT fk = m_value[f].key;
			KeyT ck = m_value[c].key;

			swap(m_value[f], m_value[c]);
			swap(m_keyIndex[fk], m_keyIndex[ck]);

			c = f;
		}

        return f;
	}
};

struct Node
{
	unsigned long long key;
	int d;
	int x;
	int y;

	Node(){d = 0;x = 0; y = 0;}
	Node(int tempx,int tempy,int tempd,unsigned long long tempkey):x(tempx),y(tempy),d(tempd),key(tempkey)
	{

	}

	bool operator<(const Node &otherNode)
	{
		return d < otherNode.d;
	}

	bool operator <= (const Node &otherNode)
	{
		return d <= otherNode.d;
	}
};

class Solution {
public:
    bool isEscapePossible(vector<vector<int> >& blocked, vector<int>& source, vector<int>& target)
    {
    	bool bRet = false;
    	int tx = target[0];
		int ty = target[1];

		unsigned long long key = 0;
		for (int i = 0; i < blocked.size(); ++i)
		{
			vector<int> &vec = blocked[i];
			setBlock.insert(GetKey(vec[0],vec[1]));
		}

		key = GetKey(source[0],source[1]);
		Node *pNode = new Node(source[0],source[1],GetDist(source[0],source[1],tx,ty),key);
		m_openQueue.Push(*pNode);
		setPathed.insert(key);

		int dx[4] = {0,0,-1,1};
		int dy[4] = {1,-1,0,0};
		while(!m_openQueue.Empty())
		{
			Node stNode = m_openQueue.Pop();
			int cx = stNode.x;
			int cy = stNode.y;
			for(int i = 0 ; i < 4 ;i++)
			{
				int nx = cx + dx[i];
				int ny = cy + dy[i];
				if( nx == tx && ny == ty)
				{
					return true;
				}

				key = GetKey(nx,ny);
				set<unsigned long long >::iterator setIter = setPathed.find(key);
				if(setIter != setPathed.end())
				{
					continue;
				}

				set<unsigned long long >::iterator setIter1 = setBlock.find(key);
				if(setIter1 != setBlock.end())
				{
					continue;
				}

				int dist = GetDist(nx,ny,tx,ty);
				Node *pNextNode = new Node(nx,ny,dist,key);
				m_openQueue.Push(*pNextNode);
				setPathed.insert(key);
			}
		}
		return false;
    }

    void GetXY(unsigned long long key,int &x,int &y)
    {
    	y = key & 0xFFFFFFFF;
    	x = key >> 32;
    }

    unsigned long long GetKey(int x,int y)
    {
    	return (unsigned long long)x << 32 | (unsigned long long)y;
    }

    int GetDist(int sx,int sy,int tx,int ty)
    {
    	int result = 0;
    	if(sx > tx)
    	{
    		result += sx - tx;
    	}
    	else
    	{
    		result += tx - sx;
    	}

    	if(sy > ty)
    	{
    		result += sy - ty;
    	}
    	else
    	{
    		result += ty - sy;
    	}

    	return result;
    }

public:
	set<unsigned long long> setBlock;
	set<unsigned long long> setPathed;

    CIndexedBinaryHeap<unsigned long long, Node> m_openQueue;
};


int main()
{
	Solution s;

	return 0;
}