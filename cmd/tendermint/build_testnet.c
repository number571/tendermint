#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <sys/stat.h>

#define LOCBSIZE 256
#define BUFFSIZE 4096

#define TENDERMINT "go run main.go"
#define LEN(obj) sizeof((obj))/sizeof((obj)[0])

#define NODEPATH(i) nodepath, (i)
#define CONFIGFILE(i) NODEPATH(i), configfile

static char buffer[BUFFSIZE];
static char nodesid[4][LOCBSIZE];

static const char *ports_to_replace[] = {"26658", "26657", "26656"};
static const char *nodepath = "./mytestnet/node";
static const char *configfile = "/config/config.toml";

static int save_nodes_id(void);
static int replace_ports(void);
static int create_scripts(void);

int main(void) {
    int err;

    err = save_nodes_id();
    if (err != 0) {
        fprintf(stderr, "error: save_nodes_id\n");
        return 1;
    }

    err = replace_ports();
    if (err != 0) {
        fprintf(stderr, "error: replace_ports\n");
        return 2;
    }

    err = create_scripts();
    if (err != 0) {
        fprintf(stderr, "error: create_scripts\n");
        return 3;
    }

    return 0;
}

static int save_nodes_id(void) {
    FILE *ptr;

    for (int i = 0; i < LEN(nodesid); ++i) {
        snprintf(buffer, BUFFSIZE, "%s show-node-id --home %s%d", TENDERMINT, NODEPATH(i));

        ptr = popen(buffer, "r");
        if (ptr == NULL) {
            return 1;
        }
        
        fgets(nodesid[i], LOCBSIZE, ptr);
        nodesid[i][strlen(nodesid[i])-1] = '\0';

        pclose(ptr);
    }

    return 0;
}

static int replace_ports(void) {
    char locbuff[LOCBSIZE];
    FILE *ptr;
    long int pos;
    char *strptr;

    for (int i = 0; i < LEN(nodesid); ++i) {
        snprintf(locbuff, LOCBSIZE, "%s%d%s", CONFIGFILE(i));

        ptr = fopen(locbuff, "r+");
        if (ptr == NULL) {
            return 1;
        }

        pos = 0;
        while (fgets(buffer, BUFFSIZE, ptr)) {
            for (int j = 0; j < LEN(ports_to_replace); ++j) {
                if (strptr = strstr(buffer, ports_to_replace[j])) {
                    fseek(ptr, pos, SEEK_SET);
                    strptr[2] = i+'0';
                    fputs(buffer, ptr);
                    fseek(ptr, pos, SEEK_SET);
                }
            }
            pos = ftell(ptr);
        }

        fclose(ptr);
    }

    return 0;
}

static int create_scripts(void) {
    char locbuff[LOCBSIZE*2];
    FILE *ptr;

    for (int i = 0; i < LEN(nodesid); ++i) {
        int in[LEN(nodesid)-1];
        for (int j = 0, k = 0; j < LEN(nodesid); ++j) {
            if (i == j) {
                continue;
            }
            in[k++] = j;
        }

        snprintf(buffer, BUFFSIZE, "%s start --home %s%d --proxy_app=kvstore --p2p.persistent_peers=\"",
            TENDERMINT, NODEPATH(i));
        for (int j = 0; j < LEN(nodesid)-1; ++j) {
            snprintf(locbuff, LOCBSIZE*2, "%s@127.0.0.1:26%d56,", nodesid[in[j]], in[j]);
            strncat(buffer, locbuff, BUFFSIZE);
        }
        buffer[strlen(buffer)-1] = '\"';

        printf("%s\n", buffer);

        snprintf(locbuff, LOCBSIZE, "run_node%d.sh", i);

        ptr = fopen(locbuff, "w");
        if (ptr == NULL) {
            return 1;
        }
        fprintf(ptr, "%s\n\n%s\n", "#!/bin/bash", buffer);
        fclose(ptr);

        chmod(locbuff, 0744);
    }

    return 0;
}
