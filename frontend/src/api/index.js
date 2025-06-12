import { users, issues } from '@/data/mockData.js'

const isDev = import.meta.env.VITE_MODE === 'DEV'
const baseUrl = !isDev ? import.meta.env.VITE_API_BASE_URL : ''

const apiService = {
  async getUsers() {
    return users
  },

  async getIssuesList() {
    if (isDev) {
      return issues
    } else {
      const response = await fetch(`${baseUrl}/issues`, {
        method: "GET",
      })
      return response.json()
    }
  },

  // TODO : 단건 조회, POST, PATCH
  async getIssueDetails(id) {
    console.log('get issue details : ', id)
    if (isDev) {
      return issues.find(issue => issue.id === parseInt(id))
    } else {
      const response = await fetch(`${baseUrl}/issue/${id}`, {
        method: "GET",
      })
      return response.json()
    }
  },

  async createIssue(issue) {
    console.log('get issue details : ', issue)
    if (isDev) {
      issues.push(issue)
      return issue // dev
    } else {
      const response = await fetch(`${baseUrl}/issue/${issue.id}`, {
        method: "POST",
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(issue)
      })

      return response.json()
    }
  },

  async updateIssue(issue) {
    console.log('update issue : ', issue)
    if(isDev) {

    } else {
      const response = await fetch(`${baseUrl}/issue/${issue.id}`, {})
    }
  }
}

export default apiService
