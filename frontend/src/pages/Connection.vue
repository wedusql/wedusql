<script lang="ts" setup>
import { ref } from "vue";
import { Connect } from "../../wailsjs/go/main/App.js";
import { notification, message } from "ant-design-vue";
import { useRouter } from "vue-router";

const router = useRouter();
const connectionType = ref('mysql');
const dsn = ref('root:123123@localhost:3306/mysql');

const filterOption = (input: string, option: any) => {
  return option.label.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};
const options = [
  { value: "pgx", label: "PostgreSQL" },
  { value: "mysql", label: "MySQL / MariaDB" },
  { value: "sqlite3_extended", label: "SQLite" },
];

const connect = async () => {
  try {
    await Connect(connectionType.value, dsn.value);
    message.success("connected");
    router.push("/query");
  } catch (err: any) {
    notification.open({
      message: "Create connection failed",
      description: err,
    });
  }
};
</script>

<template>
  <div class="flex p-3">
    <a-input-group compact>
      <a-select
        v-model:value="connectionType"
        show-search
        :options="options"
        style="width: 20%"
        :filter-option="filterOption"
        placeholder="Connection Type"
      />
      <a-input v-model:value="dsn" style="width: 60%" placeholder="Connection String" />
      <a-button type="primary" @click="connect()">Submit</a-button>
    </a-input-group>
  </div>
</template>
