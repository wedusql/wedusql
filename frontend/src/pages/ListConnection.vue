<script lang="ts" setup>
import { ref } from 'vue'
import Card from 'primevue/card';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { useToast } from "primevue/usetoast";
import { useRouter } from "vue-router";
import { TestConnection,Connect } from "../../wailsjs/go/connections/Connection.js";

const toast = useToast();
const router = useRouter();

const databases =ref([
    {
        name: "Mysql",
        url: "https://www.mysql.com/common/logos/logo-mysql-170x115.png"
    },{
        name: "Postgres",
        url: "https://www.postgresql.org/media/img/about/press/elephant.png"
    },{
        name: "Sqlite",
        url: "https://www.sqlite.org/images/sqlite370_banner.gif"
    }
])

const visible = ref(false)
const dbType = ref("")
const dbConfig = {
    name: "",
    host: "",
    port: "",
    user: "",
    password: "",
    connectionName: ""
}

const showDialog = (db: any) => {
    visible.value = true
    dbType.value = db.name
}

const checkConnection = async() => {
    try {
        let dsn = ""
        if (dbType.value === "Mysql") {
            dsn = `${dbConfig.user}:${dbConfig.password}@${dbConfig.host}:${dbConfig.port}/${dbConfig.name}`
        } else if (dbType.value === "Postgres") {
            dsn = `postgres://${dbConfig.user}:${dbConfig.password}@${dbConfig.host}:${dbConfig.port}`
        } else if (dbType.value === "Sqlite") {
            dsn = dbConfig.name
        }

        await TestConnection(dbType.value, dsn);
        toast.add({ severity: "success", summary: "Connection Success", life: 3000 });
        // router.push("/home");
      } catch (err: any) {
        toast.add({
          severity: "error",
          summary: "Connection failed",
          detail: err,
          life: 5000,
        });
      }
}

const connect = async () => {
  try {
    let dsn = ""
        if (dbType.value === "Mysql") {
            dsn = `${dbConfig.user}:${dbConfig.password}@${dbConfig.host}:${dbConfig.port}/${dbConfig.name}`
        } else if (dbType.value === "Postgres") {
            dsn = `postgres://${dbConfig.user}:${dbConfig.password}@${dbConfig.host}:${dbConfig.port}`
        } else if (dbType.value === "Sqlite") {
            dsn = dbConfig.name
        }

    await Connect(dbType.value, dsn);
    toast.add({ severity: "success", summary: "Connected", life: 3000 });
    router.push("/home");
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "Create connection failed",
      detail: err,
      life: 5000,
    });
  }
};
</script>

<template>
    <div class="flex p-3">
        <div v-for="db in databases" class="p-2">
            <Card  style="width: 25em" @click=showDialog(db)>
                <template #header>
                    <img alt="user header" :src=db.url />
                </template>
                <template #title class="center"> {{db.name}} </template>
            </Card>
            
        </div>

        <Dialog v-model:visible="visible" modal header="Configuration" :style="{ width: '50vw' }" >
            <div v-if="dbType === 'Mysql' || dbType === 'Postgress'">
                <InputText v-model="dbConfig.connectionName" type="text" placeholder="Connection Name" />
                <InputText v-model="dbConfig.host" type="text" placeholder="Host" />
                <InputText v-model="dbConfig.port" type="text" placeholder="Port" />
                <InputText v-model="dbConfig.user" type="text" placeholder="User" />
                <InputText v-model="dbConfig.password" type="password" placeholder="Passwrord" />
                <InputText v-model="dbConfig.name" type="text" placeholder="database name" />
            </div>
            <div class="flex">
                <Button  label="Test" severity="success" @click="checkConnection" />
                <Button  label="Connect" @click="connect"/>
            </div>

        </Dialog>
    </div>
</template>

<style scoped>
img {
    width: 20em;
    height: 13em;
}

.center {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>