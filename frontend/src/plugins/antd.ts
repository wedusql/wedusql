import { App } from "vue";

import "ant-design-vue/dist/antd.css";
import "ant-design-vue/es/message/style/css";
import { Button, Input, Row, Col, Table, Select } from "ant-design-vue";

export default {
  install: (app: App) => {
    app.use(Button);
    app.use(Input);
    app.use(Select);
    app.use(Row);
    app.use(Col);
    app.use(Table);
  },
};
