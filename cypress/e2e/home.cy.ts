describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/home')
    cy.intercept({
      method: 'POST',
      url: 'http://localhost:9000/walmart',
    }).as('apiCheck')
    cy.get('input').type('test')
    cy.get('#share').click()
    cy
    .wait('@apiCheck')
    .then(intercept => {
     
    });
  })
})