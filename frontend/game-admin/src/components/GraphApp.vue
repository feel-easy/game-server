<template>
  <div>
    <vue-dag :nodes="nodes" :edges="edges"></vue-dag>
    <button @click="addNode">添加节点</button>
    <button @click="addEdge">添加边</button>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { VueDag } from 'vue-dag';

export default {
  name: 'GraphApp',
  components: {
    VueDag,
  },
  setup() {
    const nodes = ref([]);
    const edges = ref([]);

    // 获取节点和边数据
    const fetchData = async () => {
      const nodesResponse = await axios.get('http://localhost:3000/api/nodes');
      const edgesResponse = await axios.get('http://localhost:3000/api/edges');
      nodes.value = nodesResponse.data;
      edges.value = edgesResponse.data;
    };

    // 添加节点
    const addNode = async () => {
      const newNode = { name: 'New Node' };
      const response = await axios.post('http://localhost:3000/api/nodes', newNode);
      nodes.value.push(response.data);
    };

    // 添加边
    const addEdge = async () => {
      if (nodes.value.length > 1) {
        const newEdge = {
          from: nodes.value[0].id,
          to: nodes.value[1].id,
        };
        const response = await axios.post('http://localhost:3000/api/edges', newEdge);
        edges.value.push(response.data);
      }
    };

    // 初始化数据
    onMounted(() => {
      fetchData();
    });

    return {
      nodes,
      edges,
      addNode,
      addEdge,
    };
  },
};
</script>

<style>
/* 添加一些样式 */
</style>
