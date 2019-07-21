Updated 7/21/2019

First, if you don't have Postman installed, get it at: https://www.getpostman.com/downloads/

Whether you are running tests or creating tests, make sure a local version of the backend is running

To run these tests in Postman:
    1. Run the sql script clear_tables.sql in the testdata directory, then run the script for your vertical's test data
    2. Open Postman and select import at the top left
    3. Choose the json file based on which vertical you are testing
    4. Click on the arrow next to the collection name in the sidebar and click the run button
    5. Click the run button inside of Postman's collection runner

To create Postman tests for your vertical:
    1. Write a sql script to populate empty database tables with test data and save it in the testdata directory
    2. Open Postman and click collection on the pop-up, if there is no pop-up, click the new button and select collection
    3. Name your collection after your vertical and hit create
    4. Click the new button again, select request and add it to your collection
    5. Enter the URL and any Headers/Body data, then navigate to the tests tab to write your tests
        Note: There are useful snippets for testing on the right
    6. Once you have created your tests, run your tests as a collection and make sure they all pass
    7. Click the three dots next to your collection name, select export, and add the json file to this directory