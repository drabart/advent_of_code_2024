#include <bits/stdc++.h>
#include <fstream>

typedef long long ll;

using std::unordered_map;
using std::vector;
using std::string;

unordered_map<ll, unordered_map<ll, ll>> skip;
ll skip_step = 5;

ll l10(ll a) 
{
	ll b = 0;
	while(a > 0) {
		b++;
		a /= 10;
	}
	return b;
}

ll Pow10[] = {
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
};

unordered_map<ll, ll> steps(ll v, ll s, bool allow_cache) {
    if(allow_cache && skip.find(v) != skip.end()) {
        return skip[v];
    }

    unordered_map<ll, ll> r;
    unordered_map<ll, ll> rc;
    r[v] = 1;

    for(ll i=0; i<s; i++) {
        rc.clear();

        for(auto& kv:r) {
            if(kv.first == 0) {
				rc[1] += kv.second;
				continue;
			}
			ll lg = l10(kv.first);
			if(lg%2 == 0) {
				rc[kv.first/Pow10[lg/2]] += kv.second;
				rc[kv.first%Pow10[lg/2]] += kv.second;
				continue;
			}
			rc[kv.first*2024] += kv.second;
        }

        r = rc;
    }

    if(allow_cache)
        skip[v] = r;
    return r;
}

ll solve(ll v, ll n)
{
    unordered_map<ll, ll> rocks;
    unordered_map<ll, ll> rocks_2;
    rocks[v] = 1;
    while(n) {
        rocks_2.clear();

        int step = std::min(skip_step, n);
        for(auto& kv : rocks) {
            auto a = steps(kv.first, step, step == skip_step);
            for(auto& kv2 : a) {
                rocks_2[kv2.first] += kv.second * kv2.second;
            }
        }

        rocks = rocks_2;
        n -= step;
    }

    ll s = 0;
    for(auto kv:rocks) {
        s += kv.second;
    }

    return s;
}

ll solve(vector<ll>& v, ll n)
{
    ll sum = 0;
    for(auto& a : v) {
        sum += solve(a, n);
    }
    return sum;
}

void part1(vector<ll>& v)
{
    printf("Part 1 solution: %lld\n", solve(v, 25));
}

void part2(vector<ll>& v)
{
    printf("Part 2 solution: %lld\n", solve(v, 75));
}

void part3(vector<ll>& v)
{
    printf("Part 3 solution: %lld\n", solve(v, 2000));
}

int main(int argc, char** argv)
{
    std::ifstream f(argv[1]);
    string line;
    getline(f, line);
    std::istringstream strm(line);

    ll a;
    vector<ll> v;

    while(strm >> a) {
        v.push_back(a);
    }

    part1(v);
    part2(v);
    // part3(v);
}