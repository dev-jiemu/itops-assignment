<template>
  <div>
    <p>이슈 생성/상세/수정 페이지입니다.</p>
  </div>
  <div v-if="!isEdit">
     담당자 :
      <select v-model="issue.user">
        <option v-for="item in users" :value="item.id">{{item.name}}</option>
      </select>
  </div>
  <div v-if="isEdit">
    <p v-if="isEdit">ID: {{issue.id}}</p>
    <p>제목 : <input v-model="issue.title"></p>
    <p>상태 :
      <select v-model="issue.status">
        <option v-for="item in statusList" :value="item">{{item}}</option>
      </select>
    </p>
  </div>
  <div style="padding-top: 10px;">
    <button type="submit" @click="doIssue">
      {{ isEdit ? '수정하기' : '생성하기' }}
    </button>
    <button type="button" @click="goList">
      취소
    </button>
  </div>
</template>
<script setup>
import apiService from '@/api/index.js'
import { useRoute, useRouter } from 'vue-router';
import { computed, onMounted, ref } from 'vue';
import { users } from '@/data/mockData.js';

const route = useRoute();
const router = useRouter();

const issue = ref({
  id: 0,
  title: '',
  description: '',
  status: '',
  user: null,
  createdAt: '',
  updatedAt: '',
});

const statusList = [
  'PENDING', 'IN_PROGRESS', 'COMPLETED', 'CANCELLED',
];

const isEdit = computed(() => {
  return issue.value.id !== 0
})

onMounted(async () => {
  // 화면 진입시 issue 객체에 따라 판별
  let seqno = route.params.id;

  if (seqno) {
    try {
      issue.value = await apiService.getIssueDetails(seqno)
    } catch (error) {
      console.error('fail get issue : ', error)
    }
  }
});

const doIssue = async () => {
  let result

  // TODO : return goList?
  if (isEdit) {
    result = await apiService.updateIssue(issue.value)
  } else {
    result = await apiService.createIssue(issue.value)
  }
}

const goList = () => {
  router.push('/issues')
}
</script>
