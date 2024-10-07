#include <stdio.h>
#include <pthread.h>
#include <semaphore.h>
#include <unistd.h>

sem_t semaphore;

void *backgroundThread(void *vargp) {
    sem_post(&semaphore);
    printf("signalled\n");
    return NULL;
}

int main() {
    sem_init(&semaphore, 0, 1);
    pthread_t thread;
    pthread_create(&thread, NULL, backgroundThread, NULL);
    sem_wait(&semaphore);
    printf("exiting\n");
    sem_destroy(&semaphore);
    pthread_join(thread, NULL);
    return 0;
}