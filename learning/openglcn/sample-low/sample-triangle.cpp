#include <iostream>
#include <cstdio>

#include "glew.h"
#include "glut.h"

using namespace std;

void keyboard(unsigned char c, int x, int y);
void display();

int main(int argc, char* argv[])
{
    cout << "OpenGL:sample-triangle" << endl;

    glutInit(&argc, argv);
    glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGBA | GLUT_DEPTH | GLUT_STENCIL);
    glutInitWindowSize(900, 600);
    glutInitWindowPosition(100, 100);
    glutCreateWindow("sample-triangle");

    if(glewInit()) {
        cerr << "error: glew init failed!" << endl;
        exit(EXIT_FAILURE);
    }

    glViewport(0, 0, 900, 600);

    glutKeyboardFunc(keyboard);
    glutDisplayFunc(display);
    glutMainLoop();

    return 0;
}

void display()
{
    glClearColor(0.2f, 0.3f, 0.3f, 1.0f);
    glClear(GL_COLOR_BUFFER_BIT);

    glutSwapBuffers();
}

void keyboard(unsigned char c, int x, int y)
{
    printf("keyboard: c:%c x:%d y:%d\n", c, x, y);
}