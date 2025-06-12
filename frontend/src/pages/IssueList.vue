<template>
  <div>
    <p>이슈 목록 페이지입니다.</p>
  </div>
  <div>
    Status Filter
    <select v-model="status">
      <option value="">상태 선택</option>
      <option v-for="item in statusList" :value="item">{{item}}</option>
    </select>
    <input type="button" value="검색" @click="doSearch">
  </div>
  <div>
    <table>
      <thead>
      <tr>
        <th>제목</th>
        <th>상태</th>
        <th>담당자</th>
        <th>생성일</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in issues" :key="item.id">
        <td @click="goDetail(item.id)" class="clickable-cell">{{item.title}}</td>
        <td>{{item.status}}</td>
        <td>
          <template v-if="item.user">{{item.user.name}}</template>
        </td>
        <td>{{item.createdAt}}</td>
      </tr>
      <tr v-if="issues.length === 0">
        <td colspan="4">검색된 결과가 없습니다.</td>
      </tr>
      </tbody>
    </table>
  </div>
  <div style="padding-top: 10px;">
    <input type="button" value="이슈 생성" @click="createIssue">
  </div>
</template>
<script setup>
import apiService from '@/api/index.js'
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const issues = ref([])
const router = useRouter()
const status = ref('')

const statusList = [
  'PENDING', 'IN_PROGRESS', 'COMPLETED', 'CANCELLED',
];

onMounted(async () => {
  await doSearch()
})

const goDetail = (id) => {
  router.push(`/issues/${id}`);
}

const createIssue = () => {
  router.push(`/issues/new`);
}

const doSearch = async () => {
  try {
    issues.value = await apiService.getIssuesList(status.value)
  } catch (error) {
    console.error('fail issue lists : ' + error)
  }
}
</script>

<style>
.clickable-cell {
  font-weight: bold;
  text-decoration: underline;
  cursor: pointer;
}

.clickable-cell:hover {
  background-color: #f0f0f0;
}
</style>
