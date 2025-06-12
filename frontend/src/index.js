import { users, issues } from '@/data/mockData.js'

const isDev = import.meta.env.VITE_MODE === 'DEV'
const baseUrl = !isDev ? import.meta.env.VITE_API_BASE_URL : ''

const apiService = {
  async getUsers() {
    if(isDev) {
      return users
    } else {
      const response = await fetch(`${baseUrl}/api/users`)
      return response.json()
    }
  },

  async getIssues() {
    if (isDev) {
      return issues
    } else {
      const response = await fetch(`${baseUrl}/api/issues`)
      return response.json()
    }
  }
}

export default apiService
