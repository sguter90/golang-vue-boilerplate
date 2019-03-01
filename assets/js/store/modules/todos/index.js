import actions from './actions'
import mutations from './mutations'

const STORAGE_KEY = 'todos';

export default {
    namespaced: true,
    state: {
        items: JSON.parse
        (window.localStorage.getItem(STORAGE_KEY) || '[]')
    },
    actions,
    mutations,
}
;