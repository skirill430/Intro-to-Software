describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/home')
    cy.get('input').type('test')
  })
})