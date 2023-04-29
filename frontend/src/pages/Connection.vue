<script lang="ts" setup>
import { ref } from "vue";
import { SaveConnection } from "../../wailsjs/go/queries/Query.js";
import { useRouter } from "vue-router";

import InputText from "primevue/inputtext";
import Button from "primevue/button";
import Dropdown from "primevue/dropdown";
import { useToast } from "primevue/usetoast";

const router = useRouter();
const connectionType = ref("mysql");
const dsn = ref("root:123123@localhost:3306/mysql");

const options = [
  { value: "pgx", label: "PostgreSQL" },
  { value: "mysql", label: "MySQL/MariaDB" },
  { value: "sqlite3_extended", label: "SQLite" },
];

const toast = useToast();

const connect = async () => {
  try {
    await SaveConnection(connectionType.value, dsn.value);
    toast.add({ severity: "success", summary: "Connected", life: 3000 });
    router.push("/query");
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
    <Dropdown
      v-model="connectionType"
      :options="options"
      filter
      option-label="label"
      option-value="value"
      class="w-2"
    />
    <InputText
      type="text"
      v-model="dsn"
      placeholder="Connection String"
      class="w-8"
    />
    <Button @click="connect()" label="Connect" class="w-2" />
  </div>
</template>
