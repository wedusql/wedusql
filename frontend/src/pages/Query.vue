<script lang="ts" setup>
import { reactive, ref } from "vue";
import { Run } from '../../wailsjs/go/queries/Query.js';

import Textarea from "primevue/textarea";
import Button from "primevue/button";
import { useToast } from "primevue/usetoast";
import DataTable from "primevue/datatable";
import Column from "primevue/column";

const query = ref("");
let queryResult = reactive({} as any);

const toast = useToast();
const run = async () => {
  try {
    const result = await Run(query.value);
    queryResult.Columns = (result.Columns || []).map((c: string, i: number) => {
      return {
        title: c,
        dataIndex: c,
        key: c,
      };
    });
    queryResult.Rows = result.Rows;
    console.log(queryResult);
  } catch (err: any) {
    toast.add({
      severity: "error",
      summary: "Query execution failed",
      detail: err,
      life: 3000,
    });
  }
};
</script>

<template>
  <div>
    <div class="flex">
      <Textarea :rows="5" v-model="query" class="w-10" />
      <Button @click="run()" label="Run" />
    </div>
  </div>

  <div>
    <DataTable :value="queryResult.Rows || []">
      <Column
        v-for="c in queryResult.Columns || []"
        :key="c.key"
        :field="c.title"
        :header="c.title"
      />
    </DataTable>
  </div>
</template>
