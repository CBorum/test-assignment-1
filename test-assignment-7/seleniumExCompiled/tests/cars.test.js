const { Builder, By, Key, until } = require('selenium-webdriver')
const axios = require('axios')

let driver 
const timeoutMs = 500

beforeAll(async () => {
    driver = await new Builder().forBrowser('chrome').build()
    await driver.get('http://localhost:3000/')
})

afterAll(async () => {
    try {
        await driver.quit()
    } catch (e) {
        console.log('driver quit error')
    }

    try {
        await axios.get('http://localhost:3000/reset')
    } catch (e) {
        // console.log('axios error')
    }
})

test('1) Verify that data is loaded, and the DOM is constructed (Five rows in the table)', async () => {
    let tbody = await driver.findElement(By.id('tbodycars'))
    let children = await tbody.findElements(By.xpath('.//tr'))
    expect(children.length).toBe(5)
})

test('2) Write 2002 in the filter text and verify that we only see two rows', async () => {
    let filterInput = await driver.findElement(By.id('filter'))
    let tbody = await driver.findElement(By.id('tbodycars'))
    filterInput.sendKeys('2002')
    await timeout(timeoutMs)
    let children = await tbody.findElements(By.xpath('.//tr'))
    expect(children.length).toBe(2)
})

test('3) Clear the text in the filter text and verify that we have the original five rows', async () => {
    let filterInput = await driver.findElement(By.id('filter'))
    let tbody = await driver.findElement(By.id('tbodycars'))
    filterInput.sendKeys(Key.BACK_SPACE, Key.BACK_SPACE, Key.BACK_SPACE, Key.BACK_SPACE)
    await timeout(timeoutMs)
    let children = await tbody.findElements(By.xpath('.//tr'))
    expect(children.length).toBe(5)
})

test('4) Click the sort “button” for Year, and verify that the top row contains the car with id 938 and the last row the car with id = 940.', async () => {
    let sortBtn = await driver.findElement(By.id('h_year'))
    
    sortBtn.click()
    await timeout(timeoutMs)

    let firstChildId = await driver.findElement(By.xpath('//*[@id="tbodycars"]/tr[1]/td[1]'))
    let lastChildId = await driver.findElement(By.xpath('//*[@id="tbodycars"]/tr[5]/td[1]'))
    
    expect(await firstChildId.getText()).toBe('938')
    expect(await lastChildId.getText()).toBe('940')
})

test('5) Press the edit button for the car with the id 938. Change the Description to "Cool car", and save changes. Verify that the row for car with id 938 now contains "Cool car" in the Description column', async () => {
    let editBtn = await driver.findElement(By.xpath('//*[@id="tbodycars"]/tr[1]/td[8]/a[1]'))
    let descriptionInput = await driver.findElement(By.id('description'))
    let saveBtn = await driver.findElement(By.id('save'))

    editBtn.click()
    descriptionInput.clear()
    timeout(timeoutMs)

    descriptionInput.sendKeys('Cool car')
    saveBtn.click()
})

test('6) Click the new “Car Button”, and click the “Save Car” button. Verify that we have an error message with the text “All fields are required” and we still only have five rows in the all cars table.', async () => {
    let newBtn = await driver.findElement(By.id('new'))
    let saveBtn = await driver.findElement(By.id('save'))
    let errText = await driver.findElement(By.id('submiterr'))
    newBtn.click()
    saveBtn.click()
    expect(await errText.getText()).toBe('All fields are required')
})

test('7) Click the new Car Button, and add a new car. Click “Save car”, and verify that the new car was added to the table with all the other cars.', async () => {
    let yearInput = await driver.findElement(By.id('year'))
    let registeredInput = await driver.findElement(By.id('registered'))
    let makeInput = await driver.findElement(By.id('make'))
    let modelInput = await driver.findElement(By.id('model'))
    let descriptionInput = await driver.findElement(By.id('description'))
    let priceInput = await driver.findElement(By.id('price'))
    let newBtn = await driver.findElement(By.id('new'))
    let saveBtn = await driver.findElement(By.id('save'))

    newBtn.click()
    yearInput.sendKeys('2008')
    registeredInput.sendKeys('2002-5-5')
    makeInput.sendKeys('Kia')
    modelInput.sendKeys('Rio')
    descriptionInput.sendKeys('As new')
    priceInput.sendKeys('31000')
    
    await timeout(timeoutMs)
    saveBtn.click()
    await timeout(timeoutMs)

    let tbody = await driver.findElement(By.id('tbodycars'))
    let children = await tbody.findElements(By.xpath('.//tr'))
    let newCar = await driver.findElements(By.xpath('//*[@id="tbodycars"]/tr[6]/td'))

    expect(children.length).toBe(6)
    expect(await newCar[0].getText()).toBe('942')
    expect(await newCar[1].getText()).toBe('2008')
    expect(await newCar[2].getText()).toBe('5/5/2002')
    expect(await newCar[3].getText()).toBe('Kia')
    expect(await newCar[4].getText()).toBe('Rio')
    expect(await newCar[5].getText()).toBe('As new')
    expect(await newCar[6].getText()).toBe('31.000,00 kr.')
})

function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
