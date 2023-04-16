#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <cstdlib>
#include <cstdio>

using namespace std;

class Object {
public:
    virtual ~Object() = default;
};

class Base {
public:
    void show() {
        cout << "Base show: " << m_id << endl;
    }

protected:
    int m_id = 0;
};

class DirveA: public Base {
public:
    virtual ~DirveA() = default;

    void show() {
        cout << "DriveA show: " << m_id << endl;
    }
};

int main(int argc, char *argv[]) {
    //case0 no virtual-table test dynamic_cast
    Base *p0 = new DirveA;
    // DirveA *p1 = dynamic_cast<DirveA*>(p0);  //compile-err: source type is not polymorphic //p0的类型不是多态，没有虚函数表
    DirveA *p1 = static_cast<DirveA*>(p0);
    // Object *p2 = static_cast<Object*>(p0);
    
    p0->show();
    p1->show();
    printf("p0:%p p1:%p \n", p0, p1);


    return 0;
}