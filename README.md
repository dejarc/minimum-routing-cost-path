# minimum-routing-cost-path

## Running against training problems
in order to calculate the average cost and runtime against the provided training data, first change the permissions of the shell script `buildAndTest.sh`. Paste the command below into a terminal in the project directory.

```
chmod +rx buildAndTest.sh
```
After permissions are updated, paste the command
```
./buildAndTest.sh
```
into the same terminal.
This will generate an output file on the host machine, `minimum-cost-routing-path`, that is then run against all tests in the `trainingProblems` folder.
