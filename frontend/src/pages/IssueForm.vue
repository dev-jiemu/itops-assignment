<template>
  <div>
    <p>이슈 생성/상세/수정 페이지입니다.</p>
  </div>
  <div>
  </div>
  <div>
    <button type="submit">
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
  // id 가 없으면 create 로 간주
  // id 가 있으면 update 로 간주함
  let result

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
