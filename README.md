# Gopher Dork
### Gopher Dork lets you create google dork queries, run them and manage the results
Developed for defensive rather than offensive security approach, Gopher Dork lets cyberteams easily manage the process of finding and attempting to remove any unwanted leaked documents from the Internet. 
The idea is to create one or more dork queries to observe and then periodically (once a day perhaps) run them. The unique results are gathered and saved in the database and the UI lets the users manage them (mark as "in progress" or "complete", etc.)

# 🔌 Requirements
- *Go 1.18+*
- *Docker*

# 🐳 How to start Gopher Dork in Docker?
1. `git clone https://github.com/mari0x00/gopher-dork.git` (clone the repo),
2. Rename the `.env.template` file to `.env` and fill out the missing fields,
> Try not to use special characters in the .env file, as some of them (e.g., `$` are not parsed correctly and might cause DB connection issues). Also, don't change the `SERVER_ADDRESS` variable, as there's a reverse proxy container that uses its' hardcoded value.
3. Run `docker-compose up -d --build`,
4. Access the web interface at `http://localhost/`.

# 🏗 How to use Gopher Dork?
1. Navigate to `Configure` tab,
2. Create your first dork: (e.g., `ext:(doc | docx | pdf | xls | xlsx | txt | ps | rtf | odt | sxw | psw | ppt | pps | xml | ppt | pptx) (intext:"Internal - FakeCompany" | intext:"Confidential - FakeCompany"`). You can limit the maximum number of results by using the `Limit` input field (providing 0 means no limit),
3. To run a single query use the `Run` button next to it, alternativelly you can run all queries using the `Run` button in the navbar,
4. Once the app is done running all the queries, you will be redirected to the `Results` page, where you can view the enties and triage them.

# 📸 Image of the UI
![image](https://github.com/mari0x00/gopher-dork/assets/25896006/f03aecd8-73f8-4f34-8b5a-d583c0abd7b3)
