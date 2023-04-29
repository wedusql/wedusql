<script lang="ts" setup>
import { reactive, ref } from "vue";
import { notification } from 'ant-design-vue';
import { Run } from '../../wailsjs/go/queries/Query.js';

const query = ref('');
let queryResult = reactive({} as any);

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
    notification.open({
      message: "Query execution failed",
      description: err,
    });
  }
};
</script>

<template>
  <div>
    <a-row>
      <a-col :span="20">
        <a-textarea :rows="5" v-model:value="query"></a-textarea>
      </a-col>
      <a-col>
        <a-button @click="run()">Run</a-button>
      </a-col>
    </a-row>

    <a-table :columns="queryResult.Columns || []" :dataSource="queryResult.Rows || []" :pagination="false">

    </a-table>

  </div>
</template>
