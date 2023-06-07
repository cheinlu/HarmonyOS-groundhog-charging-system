import Mock from 'mockjs'
import charge from './charge.json'

Mock.mock('/mock/charge','get',charge)
