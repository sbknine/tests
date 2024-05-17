import {Taxi, adjustFare} from './taxi'

describe('Taxi', () => {
  it('dist: 5,  waitTime: 5', () => {
    const result = Taxi(5, 5)
    expect(result).toEqual(25)
  }),
  it('dist: 5.4, waitTime: 5', () => {
    const result = Taxi(5.4, 5)
    expect(result).toEqual(27)
  }),
  it('dist: 5.4, waitTime: 2.49', () => {
    const result = Taxi(5.4, 2.49)
    expect(result).toEqual(25)
  }),
  it('dist: 4.7, waitTime: 2.49', () => {
    const result = Taxi(4.7, 2.49)
    expect(result).toEqual(23)
  }),
  it('dist: 5,  waitTime: 4.8', () => {
    const result = Taxi(5, 4.8)
    expect(result).toEqual(25)
  }),
  it('dist: 5,  waitTime: 1.51', () => {
    const result = Taxi(5, 1.51)
    expect(result).toEqual(22)
  })
})

describe('Min afre', () => [
  it('below minfare', () => {
    const result = adjustFare(Taxi(5, 5))
    expect(result).toEqual(35)
  }),
  it('more than minfare', () => {
    const result = adjustFare(Taxi(8, 4))
    expect(result).toEqual(36)
  })
])
