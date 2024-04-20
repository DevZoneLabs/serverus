## Branching

The idea is to try to keep all of our code clean as we progress in this project as things can get pretty complicated. 

This project has two "big" branches.
1. main: Code that is currently running in production
2. develop: Code that has not been released yet.

### Working on a new feature

It is important to set the entry point of your feature the current state of the develop branch.

Let's assume that we are working on the following feature: add-new-endpoint-to-server
1. Ensure you are starting at the develop branch level and your branch is up to date with `origin/develop`
    ```
    git checkout develop
    ```
    ```
    git fetch
    ```
    ```
    git pull
    ```
2. Create a new branch with a branch name related to the feature you are going to be working on.
    ```
    git checkout -b add-new-endpoint-to-server
    ```
3. Start your development here.
4. As you make changes you can push changes to your origin/your branch
    ```
    git push
    ```
5. It is a good practice to make your branch listen to origin/develop so it can pick up the latest changes. 
    ```
    git fetch
    git pull
    ```

6. Once you are done with your feature you can head up to github.com and open up a pull request in which you try to merge your branch to develop. Your changes will be reviewed by a team member and merged if approved. If suggestions are made you can always make changes and keep pushing them to the branch. 
7. Once the branch is merged you can delete it :).



