#include <cstdio>
#include <cstdlib>

#include <iostream>
#include <memory> //smartptr

using namespace std;

class FooA {
public:
    FooA() { cout << "FooA create:" << this << endl; }
    ~FooA() { cout << "FooA delete:" << this << endl; }
};

int main(int argc, char *argv[]) {
    {
        FooA f0;
    }

    {
        auto ptr0 = make_shared<FooA>();
    }

    return 0;
}
