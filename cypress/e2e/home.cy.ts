describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://127.0.0.1:4200/home')
    cy.intercept({
      method: 'POST',
      url: 'http://127.0.0.1:9000/walmart',
    }).as('apiCheck')
    cy.get('input').type('test')
    cy.get('#share').click()
    cy
    .wait('@apiCheck')
    .then(intercept => {
     
    });
  })
})
