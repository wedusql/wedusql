<script lang="ts" setup>
import { onMounted, reactive, watchEffect } from 'vue';
import DataTable from "primevue/datatable";
import Column from "primevue/column";

import { Run } from "../../wailsjs/go/queries/Query.js";

const props = defineProps<{
  tableName: string;
}>();

const queryResult = reactive({} as any);

const getAllTableRows = async (tableName: string) => {
  const result = await Run(`DESCRIBE ${tableName}`);
  queryResult.Columns = (result.Columns || []).map((c: string, i: number) => {
    return {
      title: c,
      dataIndex: c,
      key: c,
    };
  });
  queryResult.Rows = result.Rows;
};

watchEffect(() => {
  getAllTableRows(props.tableName);
});
</script>

<template>
  <DataTable :value="queryResult.Rows || []">
    <Column
      v-for="c in queryResult.Columns || []"
      :key="c.key"
      :field="c.title"
      :header="c.title"
    />
  </DataTable>
</template>
